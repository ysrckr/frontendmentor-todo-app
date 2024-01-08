<script lang="ts">
	import cn from 'classnames';
	import type { Todo } from '$lib/types/todo';
	import { flip } from 'svelte/animate';
	import Attribution from '../components/attribution.svelte';
	import Count from '../components/count.svelte';
	import Filter from '../components/filter.svelte';
	import Tooltip from '../components/tooltip.svelte';
	import { filter } from '$lib/filter';
	import { mode, toggleMode } from '$lib/mode';
	import type { FilterOption } from '$lib/types/filter';
	import { onMount } from 'svelte';

	let formData = '';

	const todosEndPoint = 'https://frontendmentor-todo-app.onrender.com/' + "todos"

	const controller = new AbortController();
	const signal = controller.signal;

	export let todos: Todo[] = [];

	let errorMessage;

	const getAllTodos = async () => {
		try {
			const response = await fetch(todosEndPoint, { signal });
			const data = await response.json();
			todos = [...data];
		} catch (error) {
			errorMessage = error;
		}
	};

	onMount(getAllTodos);

	const deleteTodo = async (id: number) => {
		try {
			const response = await fetch(todosEndPoint + `/${id}`, {
				method: 'DELETE',
				signal,
			});
			if (response.ok) {
				getAllTodos();
			}
		} catch (error) {
			errorMessage = error;
		}
	};

	const toggleTodoStatus = async (id: number) => {
		try {
			const response = await fetch(todosEndPoint + `/${id}`, {
				method: 'PATCH',
				signal,
			});

			if (response.ok) {
				getAllTodos();
			}
		} catch (error) {
			errorMessage = error;
		}
	};

	const filterTodos = (todos: Todo[], option: FilterOption) => {
		switch (option) {
			case 'active':
				return todos.filter((todo) => !todo.completed);
			case 'completed':
				return todos.filter((todo) => todo.completed);
			default:
				return todos;
		}
	};

	const onSubmit = async () => {
		const todo = {
			text: formData,
			completed: false,
		};

		try {
			const response = await fetch(todosEndPoint, {
				method: 'POST',
				body: JSON.stringify(todo),
				signal,
			});

			if (response.ok) {
				formData = '';
				getAllTodos();
			}
		} catch (error) {
			errorMessage = error;
		}
	};

	$: filteredTodos = filterTodos(todos, $filter);

	let hovering: any = false;

	const drop = (event: DragEvent, target: number) => {
		if (!event.dataTransfer) return;

		event.dataTransfer.dropEffect = 'move';
		const start = parseInt(event.dataTransfer.getData('text/plain'));
		const newTracklist = filteredTodos;

		if (start < target) {
			newTracklist.splice(target + 1, 0, newTracklist[start]);
			newTracklist.splice(start, 1);
		} else {
			newTracklist.splice(target, 0, newTracklist[start]);
			newTracklist.splice(start + 1, 1);
		}
		filteredTodos = newTracklist;
		hovering = null;
	};

	const dragstart = (event: DragEvent, i: number) => {
		if (!event.dataTransfer) return;

		event.dataTransfer.effectAllowed = 'move';
		event.dataTransfer.dropEffect = 'move';
		const start = i;
		event.dataTransfer.setData('text/plain', String(start));
	};
</script>

<main
	class={cn('main', {
		'main--dark': $mode === 'dark',
	})}
>
	<div
		data-testid="background"
		class={cn('background', {
			'background--dark': $mode === 'dark',
		})}
	>
		<div class="header">
			<h1
				data-testid="title"
				class="header__logo"
			>
				todo
			</h1>
			<button
				on:click={toggleMode}
				class="header__toggle"
			>
				{#if $mode === 'light'}
					<img
						src="/images/icon-moon.svg"
						alt="moon"
					/>
				{:else}
					<img
						src="/images/icon-sun.svg"
						alt="sun"
					/>
				{/if}
			</button>
		</div>
		<form
			class="form"
			on:submit={onSubmit}
		>
			<label
				class="form__label"
				for="form--todo"
			></label>
			<input
				class={cn('form__input', {
					'form__input--dark': $mode === 'dark',
				})}
				type="text"
				name="todo"
				id="form--todo"
				placeholder="Create a new todo..."
				bind:value={formData}
			/>
		</form>
		<ul
			class={cn('list', {
				'list--dark': $mode === 'dark',
			})}
		>
			{#each filteredTodos as todo, index (todo.id)}
				<li
					class="todo"
					animate:flip
					draggable={true}
					on:dragstart={(event) => dragstart(event, index)}
					on:drop|preventDefault={(event) => drop(event, index)}
					on:dragover={(event) => event.preventDefault()}
					on:dragenter={() => (hovering = index)}
				>
					<button
						class={cn('todo__checkbox', {
							'todo__checkbox--checked': todo.completed,
							'todo__checkbox--dark': $mode === 'dark',
						})}
						on:click={() => toggleTodoStatus(todo.id)}
					>
					</button>
					<p
						class={cn('todo__text', {
							'todo__text--completed': todo.completed,
						})}
					>
						{todo.text}
					</p>
					<button
						class={cn('todo__delete', {
							'todo__delete--dark': $mode === 'dark',
						})}
						on:click={async () => await deleteTodo(todo.id)}
					></button>
				</li>
			{/each}
			<Count todosLength={todos.length} />
		</ul>
	</div>
	<Filter />
	<Tooltip />
	<Attribution />
</main>

<style
	lang="scss"
	global
>
	@import '../styles/global.scss';
</style>

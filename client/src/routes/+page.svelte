<script lang="ts">
	import cn from 'classnames';
	import type { Todo } from '$lib/types/todo';
  import {flip} from 'svelte/animate';
	import Attribution from '../components/attribution.svelte';
	import Count from '../components/count.svelte';
	import Filter from '../components/filter.svelte';
	import Tooltip from '../components/tooltip.svelte';


	type Mode = 'light' | 'dark';
	let mode: Mode = 'light';

	const toggleMode = () => {
		mode = mode === 'light' ? 'dark' : 'light';
	};

	export let todos: Todo[] = [
		{
			id: 1,
			text: 'Complete online JavaScript course',
			completed: true
		},
		{
			id: 2,
			text: 'Jog around the park 3x',
			completed: false
		},
		{
			id: 3,
			text: '10 minutes meditation',
			completed: false
		},
		{
			id: 4,
			text: 'Read for 1 hour',
			completed: false
		},
		{
			id: 5,
			text: 'Pick up groceries',
			completed: false
		},
		{
			id: 6,
			text: 'Complete Todo App on Frontend Mentor',
			completed: false
		}
	];

  let hovering: any = false;

  const drop = (event: DragEvent, target: number) => {
    if(!event.dataTransfer) return

    event.dataTransfer.dropEffect = 'move'; 
    console.log(event.dataTransfer.getData("text/plain"))
    const start = parseInt(event.dataTransfer.getData("text/plain"));
    const newTracklist = todos

    if (start < target) {
      newTracklist.splice(target + 1, 0, newTracklist[start]);
      newTracklist.splice(start, 1);
    } else {
      newTracklist.splice(target, 0, newTracklist[start]);
      newTracklist.splice(start + 1, 1);
    }
    todos = newTracklist
    hovering = null
  }

  const dragstart = (event:DragEvent, i: number) => {
    if(!event.dataTransfer) return

    event.dataTransfer.effectAllowed = 'move';
    event.dataTransfer.dropEffect = 'move';
    const start = i;
    event.dataTransfer.setData('text/plain', String(start));
  }
</script>

<main class="main">
	<div
		class={cn('background', {
			'background--dark': mode === 'dark'
		})}
	>
		<div class="header">
			<h1 class="header__logo">todo</h1>
			<button
				on:click={toggleMode}
				class="header__toggle"
			>
				{#if mode === 'light'}
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
		<form class="form">
			<label class="form__label" for="form--todo"></label>
			<input
				class="form__input"
				type="text"
				name="todo"
				id="form--todo"
        placeholder="Create a new todo..."
			/>
      
		</form>
		<ul class="list">
			{#each todos as todo, index (todo.id)}
      <li
      class="todo"
      animate:flip
      draggable={true}
      on:dragstart={(event) => dragstart(event, index)}
      on:drop|preventDefault={(event) => drop(event, index)}
      on:dragover={(event) => event.preventDefault()}
      on:dragenter={() => (hovering = index)}


    >
      <button class="todo__checkbox"> </button>
      <p class="todo__text">{todo.text}</p>
      <button
        ><img
          class="todo__delete"
          src="/images/icon-cross.svg"
          alt="delete"
        /></button
      >
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

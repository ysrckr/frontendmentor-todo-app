<script lang="ts">
	import type { Todo } from '$lib/types/todo';

	import { SvelteComponent } from 'svelte';

	export class TodoItem extends SvelteComponent<Todo> {}

	export let todo: Todo;
</script>

<li class="list__item">
	<button>
		<img
			class="todo__check"
			src="/images/icon-check.svg"
			alt="check"
		/>
	</button>
	<p>{todo.text}</p>
	<button
		><img
			class="todo__delete"
			src="/images/icon-cross.svg"
			alt="delete"
		/></button
	>
</li>

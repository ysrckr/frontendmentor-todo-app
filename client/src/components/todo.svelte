<script lang="ts">
	import type { Todo } from '$lib/types/todo';

	import { SvelteComponent } from 'svelte';

	export class TodoItem extends SvelteComponent<Todo> {}

	export let todo: Todo;
</script>

<li class="todo">
	<button class="todo__checkbox">

  </button>
	<p class="todo__text">{todo.text}</p>
	<button
		><img
			class="todo__delete"
			src="/images/icon-cross.svg"
			alt="delete"
		/></button
	>
</li>

<style lang="scss">

	button {
		display: inline-block;
		background: transparent;
		border: none;
		cursor: pointer;
	}



	.todo {
    text-align: left;
    width: clamp(300px, 70%, 500px);
    display: grid;
    grid-template-columns: 1fr 3fr 1fr;
    place-items: center;

		border-bottom: 1px solid rgba(0, 0, 0, 0.2);

		&__text {
      text-align: left;
      font-size: 1rem;
			color: #000;
      justify-self: left;
		}

		&__delete {
			color: #000;
			font-weight: 700;
		}

    &__checkbox {
      width: 20px;
      height: 20px;
      border: 1px solid rgba(0, 0, 0, 0.2);
      border-radius: 50%;
      cursor: pointer;
    }

	
	}
</style>

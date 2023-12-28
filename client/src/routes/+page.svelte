<script lang="ts">
import cn from 'classnames'
import type { Todo } from '$lib/types/todo'
import TodoItem  from '../components/todo.svelte';
	import Attribution from '../components/attribution.svelte';

type Mode = 'light' | 'dark'
let mode: Mode = 'light'

const toggleMode = () => {
  mode = mode === 'light' ? 'dark' : 'light'
}

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
]

</script>


<main class="main">
  <div class={cn('background', {
    'background--dark': mode === 'dark'
  }
  )}>
    <div class="header">
      <h1 class="header__logo">todo</h1>
      <button on:click={toggleMode} class="header__toggle">
    {#if mode === 'light'}
      <img src="/images/icon-moon.svg" alt="moon">
    {:else}
      <img src="/images/icon-sun.svg" alt="sun">
    {/if}
      </button>
    </div>
    <form class="form">
      <label for="form--todo"></label>
      <input class="form__input" type="text" name="todo" id="form--todo">
    </form>
    <ul class="list">
    {#each todos as todo (todo.id)}
      <TodoItem {todo}/>
    {/each}
    </ul>
  </div>
  <Attribution />
</main>

<style lang="scss" global>
@import "../styles/global.scss";
</style>

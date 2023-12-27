<script lang="ts">
import cn from 'classnames'
type Mode = 'light' | 'dark'
let mode: Mode = 'light'

const toggleMode = () => {
  mode = mode === 'light' ? 'dark' : 'light'
}


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
  <li class="list__item">
    <div class="todo">
      <input class="todo__checkbox" type="checkbox" name="" id="">
      <label class="todo__label" for="">Complete online JavaScript course</label>
      <img class="todo__delete" src="/images/icon-cross.svg" alt="delete">
    </div>
  </li>
</ul>
</div>
</main>

<style lang="scss" global>
@import "../styles/global.scss";
</style>

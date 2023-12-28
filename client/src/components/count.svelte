<script lang="ts">
	export let todosLength: number = 0;
</script>

<div class="count">
	<div class="counter">
		<span class="counter__number">{todosLength}</span>
		<span class="counter__text">items left</span>
	</div>

	<div class="clear">
		<button class="clear__button">Clear Completed</button>
	</div>
</div>

<style lang="scss">
  button {
    display: inline-block;
    background: transparent;
    border: none;
    cursor: pointer;
  }
  .count{
    width: clamp(250px, 70%, 500px);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 0.5rem;
  }
</style>

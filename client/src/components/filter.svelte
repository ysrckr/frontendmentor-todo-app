<script lang="ts">
	import type { FilterOption } from '$lib/types/filter';
	import { writable } from 'svelte/store';

	const filter = writable<FilterOption>('all');

	const onFilterClick = (option: FilterOption) => {
		filter.set(option);
	};
</script>

<div class="filter">
	<button
		type="button"
		class="filter__button filter__button--active"
		on:click={() => onFilterClick('all')}
	>
		All
	</button>
	<button
		type="button"
		class="filter__button"
		on:click={() => onFilterClick('active')}>Active</button
	>
	<button
		type="button"
		class="filter__button"
		on:click={() => onFilterClick('completed')}>Completed</button
	>
</div>

<style lang="scss">
	.filter {
		width: clamp(300px, 70%, 500px);
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1rem 0.5rem;
		margin: auto;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
		border-radius: 0.5rem;
	}

	.filter__button {
		display: inline-block;
		background: transparent;
		border: none;
		cursor: pointer;
		font-size: 0.8rem;
		color: #000;
		font-weight: 700;
	}

	.filter__button--active {
		color: #3a7bfd;
	}
</style>

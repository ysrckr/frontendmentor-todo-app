<script lang="ts">
	import cn from 'classnames';
import {filter} from '$lib/filter'
	import type { FilterOption } from '$lib/types/filter';

	const onFilterClick = (option: FilterOption) => {
		filter.set(option);
	};
</script>

<div class="filter">
	<button
		type="button"
		class={cn('filter__button', {
			'filter__button--active': $filter === 'all',
		})}
		on:click={() => onFilterClick('all')}
	>
		All
	</button>
	<button
		type="button"
		class={cn('filter__button', {
			'filter__button--active': $filter === 'active',
		})}
		on:click={() => onFilterClick('active')}>Active</button
	>
	<button
		type="button"
		class={cn('filter__button', {
			'filter__button--active': $filter === 'completed',
		})}
		on:click={() => onFilterClick('completed')}>Completed</button
	>
</div>

<style lang="scss">
  @use 'sass:map';
  @use '../styles/partials//colors' as c;
  
	.filter {
		width: clamp(calc(300px - 1rem), 70%, 500px);
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
		color: map.get(c.$light-theme-clrs, 'dark-grayish-blue');
		font-weight: 700;
	}

	.filter__button--active {
		color: c.$primary
	}
</style>

<script lang="ts">
	export let tooltipText: string = 'Drag and drop to reorder list';
</script>

<p class="tooltip-text">
	{tooltipText}
</p>

<style lang="scss">
	@use 'sass:map';
	@use '../styles/partials/colors' as c;

	.tooltip-text {
		text-align: center;
		font-size: 0.8rem;
		font-family: 'Josefin Sans', sans-serif;
		color: map.get(c.$light-theme-clrs, 'dark-grayish-blue');
		margin: auto;
	}
</style>

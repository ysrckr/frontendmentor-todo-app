<script lang="ts">
	export let tooltipText: string = 'Drag and drop to reorder list';
</script>

<p class="tooltip-text">
	{tooltipText}
</p>

<style lang="scss">
  .tooltip-text {
    text-align: center;
    font-size: 0.8rem;
    color: #000;
    margin: auto;
  }
</style>

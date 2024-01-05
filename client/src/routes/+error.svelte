<script>
	import { page } from '$app/stores';
</script>

<div class="error">
	<h1>{$page.status}: {$page?.error?.message}</h1>
</div>

<style lang="scss">
	.error {
		margin: 25% auto;
		text-align: center;
	}
</style>

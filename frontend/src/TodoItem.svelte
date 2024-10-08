<script lang="ts">
	export let task: String;
	export let dueDate: Date | null;
	export let completed: boolean = false;
</script>

<li class="py-4">
	<div class="flex items-center">
		<input
			name="todo1"
			type="checkbox"
			class="h-4 w-4 text-teal-600 focus:ring-teal-500 border-gray-300 rounded"
			bind:checked={completed}
		/>
		<label for="todo1" class="ml-3 block text-gray-900">
			<span class="text-lg font-medium">{task}</span>
			{#if dueDate}
				<span class="text-sm font-light text-gray-500">Due on {dueDate.toLocaleDateString()}</span>
			{/if}
		</label>
	</div>
</li>

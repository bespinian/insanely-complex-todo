<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type Task from '$lib/task';

	export let task: Task;

	const dispatch = createEventDispatcher<{
		remove: string;
		toggle: Task;
	}>();

	function handleRemove(event: MouseEvent) {
		dispatch('remove', task.id);
		event.preventDefault();
	}

	function handleToggle(event: Event) {
		if (event.target) {
			dispatch('toggle', task);
			event.preventDefault();
		}
	}
</script>

<li class="py-4" id="todo-{task.id}">
	<div class="flex items-center">
		<input
			name="todo1"
			type="checkbox"
			class="h-4 w-4 text-teal-600 focus:ring-teal-500 border-gray-300 rounded"
			bind:checked={task.complete}
			on:change={handleToggle}
		/>
		<label for="todo1" class="ml-3 flex-1">
			<span
				class="text-lg font-medium"
				class:text-gray-400={task.complete}
				class:line-through={task.complete}>{task.name}</span
			>
		</label>
		<a on:click={handleRemove} href="#top" title="Delete">🪣</a>
	</div>
</li>

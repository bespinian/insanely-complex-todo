<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let id: string;
	export let task: string;
	export let completed: boolean = false;

	const dispatch = createEventDispatcher<{
		remove: string;
		toggle: { id: string; complete: boolean };
	}>();

	function handleRemove(event: MouseEvent) {
		dispatch('remove', id);
		event.preventDefault();
	}

	function handleToggle(event: Event) {
		if (event.target) {
			let checkbox = event.target as HTMLFormElement;
			dispatch('toggle', { id: id, complete: checkbox.checked });
			event.preventDefault();
		}
	}
</script>

<li class="py-4" id="todo-{id}">
	<div class="flex items-center">
		<input
			name="todo1"
			type="checkbox"
			class="h-4 w-4 text-teal-600 focus:ring-teal-500 border-gray-300 rounded"
			bind:checked={completed}
			on:change={handleToggle}
		/>
		<label for="todo1" class="ml-3 flex-1">
			<span
				class="text-lg font-medium"
				class:text-gray-400={completed}
				class:line-through={completed}>{task}</span
			>
		</label>
		<a on:click={handleRemove} href="#top" title="Delete">🪣</a>
	</div>
</li>

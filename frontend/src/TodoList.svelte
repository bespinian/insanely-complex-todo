<script lang="ts">
	import TodoItem from './TodoItem.svelte';

	let newItem = '';
	let tasks = [
		{
			id: crypto.randomUUID(),
			name: 'Complete project review',
			complete: false
		},
		{
			id: crypto.randomUUID(),
			name: 'Do another task',
			complete: true
		}
	];

	function addToList() {
		tasks = [...tasks, { id: crypto.randomUUID(), name: newItem, complete: false }];
		newItem = '';
	}

	function handleRemove(event: CustomEvent) {
		tasks = tasks.filter((item) => item.id !== event.detail);
	}
</script>

<form class="w-full mx-auto px-4 py-2">
	<div class="flex items-center border-b-2 border-teal-500 py-2">
		<input
			class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
			type="text"
			placeholder="Add a task"
			bind:value={newItem}
		/>
		<button
			class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
			type="button"
			on:click={addToList}
		>
			Add
		</button>
	</div>
</form>

<ul class="divide-y divide-gray-200 px-4">
	<li class="only-child:block hidden py-4">
		<div class="flex items-center">You have completed all tasks! 🎉</div>
	</li>
	{#each tasks as { id, name, complete }}
		<TodoItem {id} task={name} completed={complete} on:remove={handleRemove}></TodoItem>
	{/each}
</ul>

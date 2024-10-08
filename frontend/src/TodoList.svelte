<script lang="ts">
	import TodoItem from './TodoItem.svelte';

	let tasks = [
		{
			name: 'Complete project review',
			due_date: new Date(2024, 10, 31),
			complete: false
		},
		{
			name: 'Do another task',
			due_date: null,
			complete: true
		}
	];
</script>

<form class="w-full mx-auto px-4 py-2">
	<div class="flex items-center border-b-2 border-teal-500 py-2">
		<input
			class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
			type="text"
			placeholder="Add a task"
		/>
		<button
			class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
			type="button"
		>
			Add
		</button>
	</div>
</form>

<ul class="divide-y divide-gray-200 px-4">
	{#each tasks as { name, due_date, complete }}
		<TodoItem task={name} dueDate={due_date} completed={complete}></TodoItem>
	{/each}
</ul>

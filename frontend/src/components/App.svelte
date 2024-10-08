<script lang="ts">
	import type { ComponentEvents } from 'svelte';
	import TodoItem from './TodoItem.svelte';
	import Task from '$lib/task';
	import taskStore from '$lib/taskStore';

	export let tasks: Task[] = [];
	export let numberOfOpenTasks = 0;

	taskStore.subscribe((tasksFromStore) => {
		tasks = tasksFromStore;
		numberOfOpenTasks = tasks.filter((task: Task) => !task.complete).length;
	});

	let newItem = '';

	function addTask() {
		taskStore.add(newItem);
		newItem = '';
	}

	function handleRemove(event: ComponentEvents<TodoItem>['remove']) {
		taskStore.removeById(event.detail);
	}

	function handleToggle(event: ComponentEvents<TodoItem>['toggle']) {
		taskStore.toggle(event.detail.id, event.detail.complete);
	}
</script>

<div class="bg-white max-w-lg mx-auto shadow-lg rounded-lg mt-16">
	<h1 class="font-bold px-4 py-2">To-Do List</h1>

	<form class="w-full mx-auto px-4 py-2">
		<div class="flex items-center border-b-2 border-blue-500 py-2">
			<input
				class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
				type="text"
				placeholder="Add a task"
				bind:value={newItem}
			/>
			<button
				class="flex-shrink-0 bg-blue-500 hover:bg-blue-700 border-blue-500 hover:border-blue-700 text-sm border-4 text-white py-1 px-2 rounded"
				type="button"
				on:click={addTask}
			>
				Add
			</button>
		</div>
	</form>

	{#await taskStore.init()}
		<ul class="divide-y divide-gray-200 px-4">
			<li class="py-4 animate-pulse">
				<div class="h-4 bg-gray-300 rounded mb-2"></div>
			</li>
		</ul>
	{:then}
		<div class="px-4 py-2 flex">
			<span class="text-sm text-gray-500">
				{#if numberOfOpenTasks == 0}
					You have <strong>no</strong> open tasks 🎉!
				{:else if numberOfOpenTasks == 1}
					You have <strong>one</strong> open task.
				{:else}
					You have <strong>{numberOfOpenTasks}</strong> open tasks.
				{/if}
			</span>
		</div>

		<ul class="divide-y divide-gray-200 px-4">
			<li class="only-child:block hidden py-4">
				<div class="flex items-center">You have no tasks yet. Create one now.</div>
			</li>
			{#each tasks as { id, name, complete }}
				<TodoItem
					{id}
					task={name}
					completed={complete}
					on:toggle={handleToggle}
					on:remove={handleRemove}
				></TodoItem>
			{/each}
		</ul>
	{/await}
</div>

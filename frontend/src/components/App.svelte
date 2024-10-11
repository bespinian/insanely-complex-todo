<script lang="ts">
	import store from '$lib/task_store';
	import { onMount, type ComponentEvents } from 'svelte';
	import TodoItem from './TodoItem.svelte';
	import type Task from '$lib/task';

	let newItem = '';
	let fetchingTasks: Promise<void>;
	let tasks = store.tasks;
	let error = store.error;

	onMount(() => (fetchingTasks = store.fetch()));

	export let numberOfOpenTasks = 0;
	store.tasks.subscribe((tasksFromStore: Task[]) => {
		numberOfOpenTasks = tasksFromStore.filter((task: Task) => !task.complete).length;
	});

	function handleAdd() {
		store.insert(newItem);
		newItem = '';
	}

	function handleRemove(event: ComponentEvents<TodoItem>['remove']) {
		store.removeById(event.detail);
	}

	function handleToggle(event: ComponentEvents<TodoItem>['toggle']) {
		store.update(event.detail as Task);
	}
</script>

<div class="bg-white max-w-lg mx-auto shadow-lg rounded-lg mt-16">
	<h1 class="font-bold px-4 py-2">To-Do List</h1>

	{#if $error}
		<p class="px-4 py-2 text-red-700 bg-red-50">{$error}</p>
	{/if}

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
				on:click={handleAdd}
			>
				Add
			</button>
		</div>
	</form>

	{#await fetchingTasks}
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
			{#each $tasks as task}
				<TodoItem {task} on:toggle={handleToggle} on:remove={handleRemove}></TodoItem>
			{/each}
		</ul>
	{/await}
</div>

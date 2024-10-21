<script lang="ts">
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import { writable } from 'svelte/store';

	let endpoint: string = '/ws';

	type Event = { timestamp: Date; type: string; payload: { name: string } };
	const event_symbols: Record<string, string> = {
		create_task: '➕',
		delete_task: '❌',
		update_task: '✏️'
	};

	const notifications = writable<Event[]>([]);

	function removeOldestNotification() {
		notifications.update((notifications: Event[]) => {
			notifications.pop();
			return notifications;
		});
	}

	onMount(() => {
		const socket: WebSocket = new WebSocket(endpoint);

		socket.addEventListener('message', (event: MessageEvent) => {
			const data: Event = JSON.parse(event.data);

			notifications.update((notifications: Event[]) => {
				return [data, ...notifications];
			});
			setTimeout(removeOldestNotification, 2000);
		});
	});
</script>

<div class="absolute space-y-3 p-4 top-0 end-0 min-w-60">
	{#each [...$notifications].reverse() as event}
		<div
			class="max-w-xs bg-white border border-gray-200 rounded-xl shadow-lg"
			role="alert"
			tabindex="-1"
			out:fade
		>
			<div class="flex p-4">
				<div class="shrink-0">
					{event_symbols[event.type]}
				</div>
				<div class="ms-3">
					<p id="hs-toast-stack-toggle-update-label" class="text-sm text-gray-700">
						{event.payload.name}
					</p>
				</div>
			</div>
		</div>
	{/each}
</div>

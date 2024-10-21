import asyncio
import logging
from typing import Set

import websockets.asyncio.connection
from websockets.asyncio.server import ServerConnection, serve

connected_clients: Set[ServerConnection] = set()


def broadcast(message: str):
    websockets.asyncio.connection.broadcast(connected_clients.copy(), message)


async def start_websocket_server(host: str, port: int):
    async with serve(websocket_handler, host, port):
        await asyncio.get_running_loop().create_future()


async def websocket_handler(websocket: ServerConnection):
    logging.info(f"Websocket client connected from {websocket.remote_address[0]}")
    connected_clients.add(websocket)

    async for message in websocket:
        logging.info(f"Client sent a message: {str(message)}")

    connected_clients.remove(websocket)

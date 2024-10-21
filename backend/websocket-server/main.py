import asyncio
import logging
import os
from signal import SIGINT, SIGTERM

from app.pubsub import start_redis_subscription
from app.websockets import broadcast, start_websocket_server

logging.basicConfig(
    format="%(asctime)s %(message)s",
    level=logging.INFO,
)

host = os.getenv("WEBSOCKET_SERVER_HOST", "localhost")
port = int(os.getenv("WEBSOCKET_SERVER_PORT", 8765))
redis_url = os.getenv("WEBSOCKET_SERVER_REDIS_URL", "redis://redis")


def send_message_to_all(message: str):
    broadcast(message)


async def main():
    async with asyncio.TaskGroup() as tg:
        tg.create_task(start_websocket_server(host, port))
        tg.create_task(start_redis_subscription(redis_url, send_message_to_all))


if __name__ == "__main__":
    logging.info("🎸 Websocket server started.")

    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)
    main_task = asyncio.ensure_future(main())
    for signal in [SIGINT, SIGTERM]:
        loop.add_signal_handler(signal, main_task.cancel)

    try:
        loop.run_until_complete(main_task)
    except asyncio.exceptions.CancelledError:
        pass
    finally:
        loop.close()

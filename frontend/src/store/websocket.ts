import { writable, get } from 'svelte/store';
import { toggleModal } from './modal';
import { gameStore } from './game';
import { isWSMessage, type createAction } from '../types/websocket';

export type WSType = {
  ws: WebSocket | null,
  state: number,
}

// constants
const WS_URL = "ws://localhost:8000/v1/game";

const INITIAL_STATE: WSType = {
  ws: null,
  state: WebSocket.CLOSED
}

export const WSStore = writable(INITIAL_STATE);

export const openSocket = () => {
  let { ws } = get(WSStore);

  if (!!ws) return;

  ws = new WebSocket(WS_URL);
  
  ws.onopen = function(evt) {
    WSStore.update(() => ({
      ws,
      state: WebSocket.OPEN
    }));
    console.log("OPEN");
  }
  ws.onclose = function(evt) {
    console.log("CLOSE");
    WSStore.update(() => INITIAL_STATE);
  }
  ws.onmessage = function(evt) {
    const message = evt.data as string;
    const parsedMsg = JSON.parse(message);

    if (isWSMessage(parsedMsg)) {
      switch (parsedMsg.type) {
        case 1: {
          gameStore.update(() => parsedMsg.body);

          break;
        }
        default: {
          toggleModal({
            title: 'Error!',
            description: 'Something went wrong!',
            status: 'error'
          });
        }
      }
    }

    console.log("RESPONSE: " + evt.data);
  }
  ws.onerror = function(evt) {
    console.log("ERROR: " + evt);

    toggleModal({
      title: 'Error!',
      description: 'Something went wrong!',
      status: 'error'
    });
  }

  return;
}

export const closeSocket = () => {
  const { ws } = get(WSStore);

  if (!ws) return;

  ws.close();
}
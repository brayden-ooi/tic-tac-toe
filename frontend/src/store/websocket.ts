import { writable, get } from 'svelte/store';
import { toggleModal } from './modal';
import type { GamePayloadType } from '../types/game';
import { gameStore } from './game';

export type WSType = {
  ws: WebSocket | null,
  state: number,
}

export type WSMessageType = {
  type: number;
  body: GamePayloadType;
};

// constants
const WS_URL = "ws://localhost:8000/v1/game";
export const CREATE_GAME = () => JSON.stringify({ type: 'create', payload: '' });
export const JOIN_GAME = (id: string) => JSON.stringify({ type: 'join', payload: id });

const isWSMessage = (data: unknown): data is WSMessageType => {
  if (typeof data === 'object' && 'type' in data) {
    return true;
  }

  return false;
}

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
          const { id, state } = parsedMsg.body;

          gameStore.update(() => ({ id, state }));

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
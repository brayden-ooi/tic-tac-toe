import { get, writable } from 'svelte/store';
import { WSStore, openSocket } from './websocket';
import type { messageType } from '../types/websocket';
import { toggleModal } from './modal';

export type gameType = {
  id: string | null;
  state: Array<Array<string>> | null;
};

export const INITIAL_STATE: gameType = {
  id: null,
  state: null,
};

export const gameStore = writable(INITIAL_STATE);

const sendServer = (ws: WebSocket, message: messageType) => ws.send(JSON.stringify(message));

export function createGame() {
  openSocket();

  const unsub = WSStore.subscribe(({ ws, state }) => {
    if (!ws) {
      openSocket();

      return;
    }

    if (state === WebSocket.OPEN) {
      sendServer(ws, { type: 'create', payload: '' });
      unsub();
    }
  });
}

export function joinGame(id: string) {
  const unsub = WSStore.subscribe(({ ws, state }) => {
    if (!ws) {
      openSocket();

      return;
    }

    if (state === WebSocket.OPEN) {
      sendServer(ws, { type: 'join', payload: id });
      unsub();
    }
  });
}

export function updateGame(x: number, y: number) {
  const { ws, state } = get(WSStore);

  if (!ws || state != WebSocket.OPEN) {
    toggleModal({
      title: 'Error!',
      description: 'No game started!',
      status: 'error'
    });
  }

  sendServer(ws, { type: 'update', payload: [x, y] });
}
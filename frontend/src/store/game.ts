import { writable } from 'svelte/store';
import { CREATE_GAME, JOIN_GAME, WSStore, openSocket } from './websocket';

export type gameType = {
  id: string | null;
  state: Array<Array<string>> | null;
};

export const INITIAL_STATE: gameType = {
  id: null,
  state: null,
};

export const gameStore = writable(INITIAL_STATE);

export function createGame() {
  openSocket();

  const unsub = WSStore.subscribe(({ ws, state }) => {
    if (!ws) {
      openSocket();

      return;
    }

    if (state === WebSocket.OPEN) {
      ws.send(CREATE_GAME());
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
      ws.send(JOIN_GAME(id));
      unsub();
    }
  });
}

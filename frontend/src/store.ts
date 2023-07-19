import { writable } from 'svelte/store';

type stateType = {
  game: {
    id: string | null;
    state: Array<Array<string>> | null;
  };
  modal: {
    title: string;
    description: string;
    status: 'success' | 'info' | 'error';
    isOpen: boolean;
  }
}

export const INITIAL_STATE: stateType = {
  game: {
    id: null,
    state: null,
  },
  modal: {
    title: '',
    description: '',
    status: 'success',
    isOpen: false,
  }
}

export const store = writable(INITIAL_STATE);

export function toggleModal(modal: Omit<stateType["modal"], 'isOpen'>) {
  store.update((state) => ({ ...state, modal: { ...modal, isOpen: true } }))
}

export function toggleModalOff() {
  store.update((state) => ({ ...state, modal: INITIAL_STATE.modal }))
}

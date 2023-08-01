import { writable } from 'svelte/store';

export type modalType = {
  title: string;
  description: string;
  status: 'success' | 'info' | 'error';
  isOpen: boolean;
}

export const INITIAL_STATE = {
  title: '',
  description: '',
  status: 'success',
  isOpen: false,
}

export const modalStore = writable(INITIAL_STATE);

// modal operations
export function toggleModal(modal: Omit<modalType, 'isOpen'>) {
  modalStore.update(() => ({ ...modal, isOpen: true }));
}

export function toggleModalOff() {
  modalStore.update(() => INITIAL_STATE);
}
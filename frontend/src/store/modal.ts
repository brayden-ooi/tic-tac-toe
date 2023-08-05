import { writable } from 'svelte/store';

export type modalType = {
  title: string;
  description: string;
  status: 'success' | 'info' | 'error';
  isOpen: boolean;
  handleClose?: () => void;
}

export const INITIAL_STATE: modalType = {
  title: '',
  description: '',
  status: 'success',
  isOpen: false,
  handleClose: () => {},
}

export const modalStore = writable(INITIAL_STATE);

// modal operations
export function toggleModal(modal: Omit<modalType, 'isOpen'>) {
  modalStore.update(() => ({ handleClose: () => {}, ...modal, isOpen: true }));
}

export function toggleModalOff() {
  modalStore.update((store) => {
    store.handleClose();

    return INITIAL_STATE;
  });
}
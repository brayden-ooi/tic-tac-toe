<script lang="ts">
  import { INITIAL_STATE, modalStore, toggleModalOff } from '../store/modal';

  // modal state
  let modalState = INITIAL_STATE;

  modalStore.subscribe((store) => {
    modalState = store;
  });

  let dialogElement: HTMLDialogElement;
  
  $: if (modalState.isOpen) {
    dialogElement?.showModal();
  } else {
    dialogElement?.close();
  }
</script>

<dialog bind:this={dialogElement}>
  <p>{modalState.title}</p>
  <p>{modalState.description}</p>
  <button on:click={toggleModalOff}>Close</button>
</dialog>
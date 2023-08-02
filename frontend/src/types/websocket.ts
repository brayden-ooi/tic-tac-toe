import type { GamePayloadType } from "./game";

export type createAction = {
  type: 'create';
  payload: '';
};

export type joinAction = {
  type: 'join';
  payload: string
}

export type moveAction = {
  type: 'update';
  payload: [number, number];
}

export type messageType = createAction | joinAction | moveAction;

export type WSMessageType = {
  type: number;
  body: GamePayloadType;
};

export const isWSMessage = (data: unknown): data is WSMessageType => {
  if (typeof data === 'object' && 'type' in data) {
    return true;
  }

  return false;
}
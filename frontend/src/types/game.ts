// derived from GamePayload in handler_game.go
export type GamePayloadType = {
  id: string;
  state: Array<Array<string>>;
  status: 'playing' | 'completed';
  result: string | "stale";
}

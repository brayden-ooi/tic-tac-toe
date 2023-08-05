# Tic-Tac-Toe

This is an exploratory project on Svelte as the Front End and Golang as the Back End. The server uses `gorilla/websocket`, `channels` and `goroutines` to concurrently serve tic-tac-toe games. The server will be able to declare a winner for a game or declare a stale, and further block any attempts to update the game. The client is constructed with `vite + Svelte + TS`. There are a few pages to allow users to create or join a game, and `svelte-routing` is used to take care of managing routing and redirecting the users.

## Notes

1. Due to its exploratory nature, the Front End was constructed with only functionalities in mind and minimal efforts went to UX and website design.
2. For example, a `clipboard` copy and paste functionality could be added to improve the UX of capturing the game ID. And the overall pages and modal could use more polishing. 
3. Also, the project only worked out the happy paths to concurrently host tic-tac-toe games and determine the outcomes of the games at any given time. Im sure this will not scale well.

## Getting started
1. Start by go into the `/backend` and run `go build`. Then you will have an executable where you can then run `./tic-tac-toe`.
2. After the server is running, go to `/frontend` and run `npm install`, then run `npm run dev`.
3. You will need to run the app on two separate browsers in order to play a game.
4. On first instance, press `Create game`. You will be redirected to the `/game` page with the game ID and the board. 
5. Copy the game ID and paste it into the second instance. Then press `Join game`.
6. You will be redirected to the `/game` page with the same game ID. Then, the player with the first instance will be able to start the game by making a move. 

## To-do's
1. ID shorteners - send shortened IDs to Front End and create a map at the Back End for the real game IDs
2. Homepage and modal polishing
3. Clipboard copy and paste
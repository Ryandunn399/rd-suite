
// Import the express in typescript file
import express, { Express, Request, Response } from 'express';
 
// Initialize the express engine
const app: Express = express();
 
// Take a port 3000 for running server.
const port: number = 3000;
 
// Handling '/' Request
app.get('/', (req: Request, res: Response) => {
    res.send("TypeScript With Express");
});
 
// Server setup
app.listen(port, () => {
    console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});
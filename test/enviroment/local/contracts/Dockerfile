FROM node:14

WORKDIR /usr/src/app
COPY abis ./abis
COPY bin ./bin
COPY caller ./caller
COPY contracts ./contracts
COPY generated ./generated
COPY migrations ./migrations
COPY src ./src
COPY package-lock.json .
COPY package.json .
COPY schema.graphql .
COPY subgraph.yaml .
COPY start.sh .
COPY truffle-config.js .
RUN npm install

EXPOSE 3000

ENTRYPOINT ["npm", "start"]
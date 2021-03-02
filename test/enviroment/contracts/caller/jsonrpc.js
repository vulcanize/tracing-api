const fetch = require('node-fetch');

module.exports = function JsonRPC(host) {
  let id = 0;
  function body(method, params) {
    return {
      jsonrpc: '2.0',
      id: id++,
      method,
      params
    };
  }
  async function call(data, isBatch = false) {
    const resp = await fetch(host, {
      agent: null,
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
      body: JSON.stringify(data),
    });
    const body = await resp.json();

    if (!isBatch) {
      if (!!body.error) {
        throw new Error(body.error.message);
      }
      return body.result;
    }

    return body.map(({ result }) => result);
  }
  return {
    exec(method = "", ...params) {
      return call(body(method, params));
    },
    batch() {
      const data = [];
      return {
        push(method = "", ...params) {
          data.push(body(method, params))
        },
        exec() {
          return call(data, true);
        }
      }
    }
  }
}
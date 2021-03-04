pragma solidity >=0.5.12;

contract UintStorage {

  mapping(string => uint) private data;

  function get(string memory key) public view returns (uint){
    return data[key];
  }

  function set(string memory key, uint _value) public payable {
    data[key] = _value;
  }
}

contract BlockNumStorage {

  UintStorage public store = new UintStorage();

  function sync(string memory key) public payable returns (uint) {
    uint n = store.get(key);
    if (block.number != n){
      store.set.value(msg.value)(key, block.number);
    }
    return store.get(key);
  }

}
#!/bin/bash

failCount=0
passCount=0

failTest () {
  echo "- Test $1 failed."
  ((failCount++))
}

passTest () {
  echo "+ Test $1 passed."
  ((passCount++))
}

# Check if listTodos returns all 3 dummy data elements
testListTodosListsTodos() {
  listTodosResp=$(curl -s --request GET 'localhost:8081/')

  # Get length of listTodos response json array
  len=$(echo "$listTodosResp" | jq 'length')

  if  [[ $len -eq 3  ]]; then 
    passTest "${FUNCNAME[0]}" 
  else 
    failTest "${FUNCNAME[0]}"
  fi
}

testGetRandomTodoReturnsOneOfTodos () {
  randomTodoResp=$(curl -s --request GET 'localhost:8081/random')

  # Get id of random todo
  id=$(echo "$randomTodoResp" | jq '.id')

  # Get all ids returned by listTodos
  allIds=$(echo "$listTodosResp" | jq '.[].id')

  # Check if listTodo response contained the id returned
  if [[ $allIds =~ $id ]]; then
    passTest "${FUNCNAME[0]}"
  else
    failTest "${FUNCNAME[0]}"
  fi
}

testAddTodoAddsTodo () {
  addTodoResp=$(curl -s --header "Content-Type: application/json" \
                --request POST \
                --data '{"content":"newTODO"}' \
                'localhost:8081/add')

  # Response status should be ok
  status=$(echo "$addTodoResp" | jq '.status')

  if [[ "$status" != '"ok"' ]]; then
    failTest "${FUNCNAME[0]}"
    return
  fi

  afterAddListTodosResp=$(curl -s --request GET 'localhost:8081/')
  
  # Element at index 3 of listTodo response should be new added todo
  newTodo=$(echo "$afterAddListTodosResp" | jq '.[3].content')

  if [[ $newTodo == '"newTODO"' ]]; then
    passTest "${FUNCNAME[0]}"
  else
    failTest "${FUNCNAME[0]}"
  fi

}

testCompleteTodoCompletesTodo () {
  beforeCompleteListTodosResp=$(curl -s --request GET 'localhost:8081/')

  # Get is_completed of 4th todo in the system before complete request
  statusBeforeComplete=$(echo "$beforeCompleteListTodosResp" | jq '.[3].is_completed')

  completeTodoResp=$(curl -s --request PATCH 'localhost:8081/complete/0')

  # Response status should be ok
  status=$(echo "$completeTodoResp" | jq '.status')

  if [[ "$status" != '"ok"' ]]; then
    failTest "${FUNCNAME[0]}"
    return
  fi

  afterCompleteListTodosResp=$(curl -s --request GET 'localhost:8081/')

  # Get is_completed of 4th todo in the system after complete request
  statusAfterComplete=$(echo "$afterCompleteListTodosResp" | jq '.[3].is_completed')

  # Check if is_completed is changed from false to true after request
  if  [[ !$statusBeforeComplete && statusAfterComplete ]]; then 
    passTest "${FUNCNAME[0]}" 
  else 
    failTest "${FUNCNAME[0]}"
  fi
}

testRemoveTodoRemovesTodo () {
  beforeDeleteListTodosResp=$(curl -s --request GET 'localhost:8081/')

  # Get length of listTodos response before delete
  lenBeforeDelete=$(echo "$beforeDeleteListTodosResp" | jq 'length')

  removeTodoResp=$(curl -s --request DELETE 'localhost:8081/remove/3')

  # Response status should be ok
  status=$(echo "$removeTodoResp" | jq '.status')

  if [[ "$status" != '"ok"' ]]; then
    failTest "${FUNCNAME[0]}"
    return
  fi

  afterDeleteListTodosResp=$(curl -s --request GET 'localhost:8081/')

  # Get length of listTodos response after delete
  lenAfterDelete=$(echo "$afterDeleteListTodosResp" | jq 'length')

  # Check if result of listTodos is shorter after delete
  if  [[ $lenBeforeDelete -gt $lenAfterDelete ]]; then 
    passTest "${FUNCNAME[0]}" 
  else 
    failTest "${FUNCNAME[0]}"
  fi
}


testListTodosListsTodos
testGetRandomTodoReturnsOneOfTodos
testAddTodoAddsTodo
testCompleteTodoCompletesTodo
testRemoveTodoRemovesTodo


echo "Tests passed: $passCount"
echo "Tests failed: $failCount"

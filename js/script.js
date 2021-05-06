// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains the JS functions for index.html

function userInputClick() {
  const streetNumber = parseInt(document.getElementById("number-entered").value)
  const streetName = document.getElementById("name-entered").value
  document.getElementById('address').innerHTML = 'Your address is: ' + streetNumber + ' ' + streetName + '.'
}

/* globals Chart:false, feather:false */

(() => {
  'use strict'

  feather.replace({ 'aria-hidden': 'true' })

  // Graphs
  const ctx = document.getElementById('myChart')
  // eslint-disable-next-line no-unused-vars
  fetch('/chartdata', {
    method: 'POST',
    headers: {
      'content-type':'application/json'
    },
    body: "7d"
  })
  .then(response => response.json())
    .then(data => {
      const myChart = new Chart(ctx, {
        type: 'line',
        data: data,
        options: {
          plugins: {
            legend: {
              display: false
            },
            tooltip: {
              boxPadding: 3
            }
          }
        }
      })
    })
})()

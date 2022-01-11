var endpoints = ['a','b']

async function interateEnpoints(i){
    if(i < endpoints.length){
    	console.log('call done')
      await callEndpoint(endpoints[i])
      i = i + 1
    	interateEnpoints(i);
    }
}


async function callEndpoint(endpoint){
	let response = await fetch('https://jsfiddle.net/')
	console.log('called endpoint', response)
}

interateEnpoints(0);

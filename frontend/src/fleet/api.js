const host = 'http://localhost:8080';

async function handleFetch(fetchPromise) {
    let response;

    try {
        response = await fetchPromise;
    } catch(err) {
        throw ['error.networkFailure'];
    }

    if(!response.ok) {
        let errors = ['error.unknown'];
        try {
            errors = await response.json();
        } catch(e) {
            console.error(e)
        }
        throw errors;
    }

    try {
        return await response.json();  
    } catch(e) {
        console.error(e);
        throw ['error.invalidJson']
    }
}

export async function getAllVehicles() {
    return await handleFetch(
        fetch(host + '/vehicles')
    );
}

export async function addVehicle(vehicle) {
    return await handleFetch(
        fetch(host + '/vehicles', {
            method: 'POST',
            body: JSON.stringify(vehicle)
        })
    );
}
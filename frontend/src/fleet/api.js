const host = 'http://localhost:8080';

export async function fetchAllVehicles() {
    const response = await fetch(host + '/vehicles');
    const vehicles = await response.json();
    return vehicles;
}

export async function addVehicle(vehicle) {
    let response;
    try {
        response = await fetch(host + '/vehicles', {
            method: 'POST',
            body: JSON.stringify(vehicle)
        });
    } catch(err) {
        throw ['error.networkFailure'];
    }

    if(!response.ok) {
        let errors = ['error.unknown'];
        try {
            errors = await response.json();
        } catch(e) {
            console.error(e);
        }
        console.log(errors);
        throw errors;
    }

    return await response.json();
}
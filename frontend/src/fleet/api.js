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
            console.error(e, response.status)
        }
        throw errors;
    }

    try {
        return await response.json();  
    } catch(e) {
        console.error(e);
        throw ['error.unknown']
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

export async function getVehicleDetail(vehicleId) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}`)
    );
}

export async function getRefuellingsByVehicle(vehicleId) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}/refuellings`)
    );
}

export async function addRefuellingToVehicle(vehicleId, refuelling) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}/refuellings`, {
            method: 'POST',
            body: JSON.stringify(refuelling)
        })
    )
}

export function formToJson(form) {
    let obj = {};
    const formData = new FormData(event.target);
    formData.forEach((value, key) => {
        obj[key] = value;
    });
    return obj;
}
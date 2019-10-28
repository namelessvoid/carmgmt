const host = 'http://localhost:8080';

export async function handleFetch(fetchPromise) {
    let response;

    try {
        response = await fetchPromise;
    } catch(err) {
        console.error(err);
        throw new Error('error.networkFailure');
    }

    if(!response.ok) {
        let error = 'error.unknown';
        try {
            error = await response.json();
        } catch(e) {
            console.error(e)
        }
        throw new Error(error);
    }

    try {
        return await response.json();  
    } catch(e) {
        console.error(e);
        throw new Error('error.unknown');
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

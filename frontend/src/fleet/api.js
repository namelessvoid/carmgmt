const host = 'http://localhost:8080';

import { getToken } from '../auth/auth';
import { navigate } from 'svelte-routing';

export async function handleFetch(fetchPromise) {
    let response;

    try {
        response = await fetchPromise;
    } catch(err) {
        console.error(err);
        throw new Error('error.networkFailure');
    }

    if(response.status === 401 || response.status === 403) {
        navigate("/");
        return null;
    } else if(!response.ok) {
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
        fetch(host + '/vehicles', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${await getToken()}`
            }
        })
    );
}

export async function addVehicle(vehicle) {
    return await handleFetch(
        fetch(host + '/vehicles', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${await getToken()}`
            },
            body: JSON.stringify(vehicle)
        })
    );
}

export async function getVehicleDetail(vehicleId) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${await getToken()}`
            }
        })
    );
}

export async function getRefuellingsByVehicle(vehicleId) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}/refuellings`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${await getToken()}`
            }
        })
    );
}

export async function addRefuellingToVehicle(vehicleId, refuelling) {
    return await handleFetch(
        fetch(host + `/vehicles/${vehicleId}/refuellings`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${await getToken()}`
            },
            body: JSON.stringify(refuelling)
        })
    )
}

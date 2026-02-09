function initMap(locations) {

    console.log("Starting the map with the cities:", locations);

    if (!document.getElementById('map')) {
        console.error("Error: No div id='map' found!");
        return;
    }

    var map = L.map('map').setView([48.8566, 2.3522], 4);

    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; OpenStreetMap',
        className: 'map-tiles'
    }).addTo(map);
    async function addMarker(city) {
        const query = city.replace(/-/g, ' ').replace(/_/g, ' ');

        try {
            const response = await fetch(`https://nominatim.openstreetmap.org/search?format=json&q=${query}&limit=1`);
            const data = await response.json();

            if (data.length > 0) {
                const lat = data[0].lat;
                const lon = data[0].lon;

                L.marker([lat, lon]).addTo(map)
                    .bindPopup(`<b style="color:black">${query.toUpperCase()}</b>`);
            } else {
                console.warn("City not found :", city);
            }
        } catch (error) {
            console.error("API error :", error);
        }
    }
    if (locations && locations.length > 0) {
        locations.forEach((city, index) => {
            setTimeout(() => {
                addMarker(city);
            }, index * 1000);
        });
    }
}
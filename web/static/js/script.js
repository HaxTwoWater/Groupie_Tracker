function normalizeLocation(loc) {
    return String(loc || "").replaceAll("_", " ").replaceAll("-", " ").trim();
}

async function geocodeOSM(query) {
    const url = `https://nominatim.openstreetmap.org/search?format=json&limit=1&q=${encodeURIComponent(query)}`;
    const res = await fetch(url);
    if (!res.ok) return null;
    const data = await res.json();
    if (!Array.isArray(data) || data.length === 0) return null;
    return { lat: parseFloat(data[0].lat), lon: parseFloat(data[0].lon) };
}

function initMap(markersData) {
    var map = L.map('map').setView([48.8566, 2.3522], 4);

    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; OpenStreetMap',
        className: 'map-tiles'
    }).addTo(map);

    async function addMarker(item) {
        const query = normalizeLocation(item.location);
        const geo = await geocodeOSM(query);
        if (!geo) return;

        const dates = (item.dates || []).slice(0, 6).join("<br>");
        const popup = `
      <div style="color:black">
        <b>${item.artist}</b><br>
        <span>${query}</span><br><br>
        <b>Dates :</b><br>${dates}
      </div>
    `;

        L.marker([geo.lat, geo.lon]).addTo(map).bindPopup(popup);
    }

    (markersData || []).slice(0, 60).forEach((item, i) => {
        setTimeout(() => addMarker(item), i * 800);
    });
}

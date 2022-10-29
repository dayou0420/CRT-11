async function myFetch(url) {
    const response = await fetch(url);
    const json = await response.json();
    return json;
}

(async function() {
    const f = await myFetch('https://crt-11.onrender.com/v1/tasks');
    const date = f.data.data.map(m => m.date);
    const bill = f.data.data.map(m => m.bill)

    const data = {
        labels: date,
        datasets: [{
            label: 'Utility cost',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderColor: 'gba(75, 192, 192, 1)',
            data: bill,
        }]
    };
    const config = {
        type: 'line',
        data: data,
        options: {
            layout: {
                padding: 20
            }
        }
    };

    const myChart = new Chart(
        document.getElementById('myChart'),
        config
    );
})();

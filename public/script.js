const DEV_URL = 'https://crt-11.onrender.com/v1/tasks'
const LOCAL_URL = 'http://localhost:8080/v1/tasks'

async function myFetch(url) {
    const response = await fetch(url);
    const json = await response.json();
    return json;
}

(async function() {
    const f = await myFetch(DEV_URL);
    const date = f.data.data.map(m => m.date);
    const powerBill = f.data.data.map(m => m.power.bill);
    const powerUsed = f.data.data.map(m => m.power.used);
    const gasBill = f.data.data.map(m => m.gas.bill);
    const gasUsed = f.data.data.map(m => m.gas.used);
    const data = {
        labels: date,
        datasets: [
            {
                label: 'Power bill',
                backgroundColor: 'rgba(255, 206, 86, 0.2)',
                borderColor: 'rgba(255, 206, 86, 1)',
                data: powerBill,
            },
            {
                label: 'Power used',
                backgroundColor: 'rgba(255, 159, 64, 0.2)',
                borderColor: 'rgba(255, 159, 64, 1)',
                data: powerUsed,
            },
            {
                label: 'Gas bill',
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                data: gasBill,
            },
            {
                label: 'Gas used',
                backgroundColor: 'rgba(153, 102, 255, 0.2)',
                borderColor: 'rgba(153, 102, 255, 1)',
                data: gasUsed,
            }
        ]
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

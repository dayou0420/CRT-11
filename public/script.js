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
    const bill = f.data.data.map(m => m.bill);
    const powerBill = f.data.data.map(m => m.power.bill);
    const powerUsed = f.data.data.map(m => m.power.used);
    const gasBill = f.data.data.map(m => m.gas.bill);
    const gasUsed = f.data.data.map(m => m.gas.used);
    const lineData = {
        labels: date,
        datasets: [
            {
                label: 'Gas and power bill',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderColor: 'rgba(75, 192, 192, 1)',
                data: bill,
            },
            {
                label: 'Power bill',
                backgroundColor: 'rgba(255, 206, 86, 0.2)',
                borderColor: 'rgba(255, 206, 86, 1)',
                data: powerBill,
            },
            {
                label: 'Power used',
                backgroundColor: 'rgba(153, 102, 255, 0.2)',
                borderColor: 'rgba(153, 102, 255, 1)',
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
                backgroundColor: 'rgba(255, 159, 64, 0.2)',
                borderColor: 'rgba(255, 159, 64, 1)',
                data: gasUsed,
            }
        ]
    };
    const lineConfig = {
        type: 'line',
        data: lineData,
        options: {
            layout: {
                padding: 40
            }
        }
    };
    const myLineChart = new Chart(
        document.getElementById('myLineChart'),
        lineConfig
    );

    const powerBillSum = powerBill.reduce((pre, cur) => pre + cur, 0);
    const gasBillSum = gasBill.reduce((pre, cur) => pre + cur, 0);

    const doughnutData = {
        labels: [
            'Power bill',
            'Gas bill'
        ],
        datasets: [{
            data: [powerBillSum, gasBillSum],
            backgroundColor: [
                'rgba(255, 206, 86, 1)',
                'rgba(255, 99, 132, 1)'
        ],
            hoverOffset: 4
        }]
    };
    const doughnutConfig = {
        type: 'doughnut',
        data: doughnutData,
        options: {
            layout: {
                padding: 200
            }
        }
    };
    const myDoughnutChart = new Chart(
        document.getElementById('myDoughnutChart'),
        doughnutConfig
    );
})();

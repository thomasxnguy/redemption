import React from "react";
import { Bar, Line } from "react-chartjs-2";
import {
  Card,
  CardBody,
  CardHeader,
  CardTitle,
  Col,
  Row,
  Table
} from "reactstrap";
import { CustomTooltips } from "@coreui/coreui-plugin-chartjs-custom-tooltips";
import { getStyle, hexToRgba } from "@coreui/coreui/dist/js/coreui-utilities";

const brandSuccess = getStyle("--success");
const brandInfo = getStyle("--info");
const brandDanger = getStyle("--danger");

// Card Chart 3
const cardChartData3 = {
  labels: ["January", "February", "March", "April", "May", "June", "July"],
  datasets: [
    {
      label: "My First dataset",
      backgroundColor: "rgba(255,255,255,.2)",
      borderColor: "rgba(255,255,255,.55)",
      data: [78, 81, 80, 45, 34, 12, 40]
    }
  ]
};

const cardChartOpts3 = {
  tooltips: {
    enabled: false,
    custom: CustomTooltips
  },
  maintainAspectRatio: false,
  legend: {
    display: false
  },
  scales: {
    xAxes: [
      {
        display: false
      }
    ],
    yAxes: [
      {
        display: false
      }
    ]
  },
  elements: {
    line: {
      borderWidth: 2
    },
    point: {
      radius: 0,
      hitRadius: 10,
      hoverRadius: 4
    }
  }
};

// Card Chart 4
const cardChartData4 = {
  labels: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
  datasets: [
    {
      label: "My First dataset",
      backgroundColor: "rgba(255,255,255,.3)",
      borderColor: "transparent",
      data: [78, 81, 80, 45, 34, 12, 40, 75, 34, 89, 32, 68, 54, 72, 18, 98]
    }
  ]
};

const cardChartOpts4 = {
  tooltips: {
    enabled: false,
    custom: CustomTooltips
  },
  maintainAspectRatio: false,
  legend: {
    display: false
  },
  scales: {
    xAxes: [
      {
        display: false,
        barPercentage: 0.6
      }
    ],
    yAxes: [
      {
        display: false
      }
    ]
  }
};

// Main Chart

//Random Numbers
function random(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

var elements = 27;
var data1 = [];
var data2 = [];
var data3 = [];

for (var i = 0; i <= elements; i++) {
  data1.push(random(50, 200));
  data2.push(random(80, 100));
  data3.push(65);
}

const mainChart = {
  labels: [
    "Mo",
    "Tu",
    "We",
    "Th",
    "Fr",
    "Sa",
    "Su",
    "Mo",
    "Tu",
    "We",
    "Th",
    "Fr",
    "Sa",
    "Su",
    "Mo",
    "Tu",
    "We",
    "Th",
    "Fr",
    "Sa",
    "Su",
    "Mo",
    "Tu",
    "We",
    "Th",
    "Fr",
    "Sa",
    "Su"
  ],
  datasets: [
    {
      label: "My First dataset",
      backgroundColor: hexToRgba(brandInfo, 10),
      borderColor: brandInfo,
      pointHoverBackgroundColor: "#fff",
      borderWidth: 2,
      data: data1
    },
    {
      label: "My Second dataset",
      backgroundColor: "transparent",
      borderColor: brandSuccess,
      pointHoverBackgroundColor: "#fff",
      borderWidth: 2,
      data: data2
    },
    {
      label: "My Third dataset",
      backgroundColor: "transparent",
      borderColor: brandDanger,
      pointHoverBackgroundColor: "#fff",
      borderWidth: 1,
      borderDash: [8, 5],
      data: data3
    }
  ]
};

const mainChartOpts = {
  tooltips: {
    enabled: false,
    custom: CustomTooltips,
    intersect: true,
    mode: "index",
    position: "nearest",
    callbacks: {
      labelColor: function(tooltipItem, chart) {
        return {
          backgroundColor:
            chart.data.datasets[tooltipItem.datasetIndex].borderColor
        };
      }
    }
  },
  maintainAspectRatio: false,
  legend: {
    display: false
  },
  scales: {
    xAxes: [
      {
        gridLines: {
          drawOnChartArea: false
        }
      }
    ],
    yAxes: [
      {
        ticks: {
          beginAtZero: true,
          maxTicksLimit: 5,
          stepSize: Math.ceil(250 / 5),
          max: 250
        }
      }
    ]
  },
  elements: {
    point: {
      radius: 0,
      hitRadius: 10,
      hoverRadius: 4,
      hoverBorderWidth: 3
    }
  }
};

const Stats = props => {
  return (
    <div>
      <Row>
        <Col xs="12" sm="6" lg="3">
          <Card className="text-white bg-indigo">
            <CardBody className="pb-0">
              <div className="text-value">214</div>
              <div>Active Links</div>
            </CardBody>
            <div className="chart-wrapper" style={{ height: "70px" }}>
              <Line
                data={cardChartData3}
                options={cardChartOpts3}
                height={70}
              />
            </div>
          </Card>
        </Col>

        <Col xs="12" sm="6" lg="3">
          <Card className="text-white bg-green">
            <CardBody className="pb-0">
              <div className="text-value">18</div>
              <div>24hr Redemptions</div>
            </CardBody>
            <div className="chart-wrapper mx-3" style={{ height: "70px" }}>
              <Bar data={cardChartData4} options={cardChartOpts4} height={70} />
            </div>
          </Card>
        </Col>
        <Col xs="12" sm="6" lg="3">
          <Card className="text-white bg-blue">
            <CardBody className="pb-0">
              <div className="text-value">300</div>
              <div>Links Redeemed</div>
            </CardBody>
            <div className="chart-wrapper mx-3" style={{ height: "70px" }}>
              <Bar data={cardChartData4} options={cardChartOpts4} height={70} />
            </div>
          </Card>
        </Col>
        <Col xs="12" sm="6" lg="3">
          <Card className="text-white bg-pink">
            <CardBody className="pb-0">
              <div className="text-value">0</div>
              <div>Errors</div>
            </CardBody>
            <div className="chart-wrapper mx-3" style={{ height: "70px" }}>
              <Bar data={cardChartData4} options={cardChartOpts4} height={70} />
            </div>
          </Card>
        </Col>
      </Row>
      <Row>
        <Col>
          <Card>
            <CardBody>
              <Row>
                <Col sm="5">
                  <CardTitle className="mb-0">Redeems</CardTitle>
                  <div className="small text-muted">Last 30 days</div>
                </Col>
                <Col sm="7" className="d-none d-sm-inline-block"></Col>
              </Row>
              <div
                className="chart-wrapper"
                style={{ height: 300 + "px", marginTop: 40 + "px" }}
              >
                <Line data={mainChart} options={mainChartOpts} height={300} />
              </div>
            </CardBody>
          </Card>
        </Col>
      </Row>

      <Row>
        <Col>
          <Card>
            <CardHeader>Latest Redemptions</CardHeader>
            <CardBody>
              <Table
                hover
                responsive
                className="table-outline mb-0 d-none d-sm-table"
              >
                <thead className="thead-light">
                  <tr>
                    <th>Code</th>
                    <th>Recipient</th>
                    <th>Amount</th>
                    <th>IP</th>
                    <th className="text-center">Country</th>
                    <th>Time</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>
                      <small className="text-muted">PK2-4Eu-2eD</small>
                    </td>
                    <td>
                      <div>bnb1s9nfkwek...</div>
                      <div className="small text-muted">
                        <span>New Install</span>
                      </div>
                    </td>
                    <td>
                      <div className="small text-muted">10 BNB</div>
                      <div className="small text-muted">10 BUSD</div>
                    </td>

                    <td>
                      <small className="text-muted">192.348.283.194</small>
                    </td>
                    <td className="text-center">
                      <i
                        className="flag-icon flag-icon-us h4 mb-0"
                        title="us"
                        id="us"
                      ></i>
                    </td>
                    <td>
                      <div className="small text-muted">
                        20 Jan 2020 11:32am
                      </div>
                      <strong>10 sec ago</strong>
                    </td>
                  </tr>
                </tbody>
              </Table>
            </CardBody>
          </Card>
        </Col>
      </Row>
    </div>
  );
};

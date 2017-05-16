/**
 * Created by at15 on 5/13/17.
 */
// start of Vue
var app = new Vue({
    el: '#app',
    data: {
        info: {},
        metrics: {},
        start_time: moment().format('YYYY-MM-DDTHH:mm'),
        end_time: moment().add(1, 'h').format('YYYY-MM-DDTHH:mm')
    },
    filters: {
        timestamp: function (str) {
            // x means Unix ms timestamp
            return moment(str).format('x');
        }
    },
    methods: {
        reset: function (event) {
            var edt = ace.edit('editor');
            // TODO: keep the history so we can travel back in time
            console.log('before reset', edt.getValue());
            edt.setValue('');
        },
        readTemplate: function (event) {
          var editor = ace.edit('editor');
          var readTmpl = {
              start_time: str2Timestamp(this.start_time),
              end_time: str2Timestamp(this.end_time),
              queries: [
                {
                  name: 'cpu.idle',
                  tags: {machine: 'machine-01', os: 'ubuntu'},
                  match_policy: 'exact'
                },
                {
                  name: 'mem.free',
                  tags: {host: 'at15-PC4073'},
                  match_policy: 'exact'
                }
              ]
          };
          // TODO: example of read by filter
          editor.setValue(JSON.stringify(readTmpl, null, '\t'));
        },
        writeTemplate: function (event) {
            var editor = ace.edit('editor');
            var writeTmpl = [
              {
                name: 'cpu.idle',
                tags: {machine: 'machine-01', os: 'ubuntu'},
                points: [[moment.now(), randomInt(0, 100)]]
              },
              {
                name: 'cpu.idle',
                tags: {machine: 'machine-02', os: 'ubuntu'},
                points: [[moment.now(), randomInt(0, 100)]]
              }
            ];
            editor.setValue(JSON.stringify(writeTmpl, null, '\t'));
        },
        read: function (event) {
            console.log('need to read');
            var editor = ace.edit('editor');
            var val = editor.getValue();
            // do the ajax request to the server and grab the data
            try {
                JSON.parse(val);
            } catch (e) {
                console.error('can\'t parse read payload', e);
                return;
            }
            axios.post('/read', JSON.parse(val))
                .then(function (response) {
                    console.log(response.data);
                    if (response.data.metrics.length === 0) {
                        console.log('no metrics returned, no need to draw the graph');
                        return;
                    }
                    // FIXME: currently, only draw the first series, should figure out a better way for handling multiple series
                    // i.e. series with same name should be draw in a same graph because they have same range
                    // TODO: the graph seems to be a straight line, could be the collector's problem again
                    var chart = echarts.getInstanceByDom(document.getElementById('canvas'));
                    firstSeries =  response.data.metrics[0];
                    // remove the existing graph
                    chart.clear();
                    var option = {
                        title: {
                            text: 'stack lines'
                        },
                        xAxis: {
                            type: 'time',
                            splitLine: {
                                show: false
                            }
                        },
                        yAxis: {
                            name: 'val'
                        },
                        series: [
                            {
                                name: firstSeries.name,
                                type: 'line',
                                // TODO: do I need to change timestamp into other format? Nop
                                data: firstSeries.points
                            }
                        ]
                    };
                    chart.setOption(option);
                })
                .catch(function (error) {
                    console.error(error);
                    console.log(error.response);
                });
        },
        write: function (event) {
            console.log('need to write');
            var editor = ace.edit('editor');
            var val = editor.getValue();
            // do the ajax request to the server and grab the data
            try {
                JSON.parse(val);
            } catch (e) {
                console.error('can\'t parse write payload', e);
                return;
            }
            axios.post('/write', JSON.parse(val))
                .then(function (response) {
                    console.log(response);
                    console.log(response.data);
                })
                .catch(function (error) {
                    console.error(error);
                    console.log(error.response);
                });
        },
        refreshInfo: function (event) {
            console.log('need to refresh database info');
            var self = this;
            axios.get('/info')
                .then(function (response) {
                    // axios will parse JSON automatically
                    self.info = response.data;
                })
                .catch(function (error) {
                    console.error(error);
                });
        }
    },
    mounted: function () {
        var self = this;
        self.refreshInfo();
    }
});
// end of Vue

// start of ACE editor
var editor = ace.edit("editor");
editor.setTheme("ace/theme/monokai");
editor.getSession().setMode("ace/mode/json");
editor.$blockScrolling = Infinity;
document.getElementById('editor').style.fontSize = '18px';
// end of ACE editor

// start of echarts
var onlyChart = echarts.init(document.getElementById('canvas'));

var option = {
    title: {
        text: 'stack lines'
    },
    tooltip: {},
    toolbox: {
        feature: {
            dataView: {show: true, readOnly: true},
            magicType: {show: true, type: ['line', 'bar']},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    legend: {
        data: ['Xephon-K(Mem)', 'Xephon-K(Cassandra)', 'KairosDB', 'InfluxDB'],
        //            orient: 'vertical',
        orient: 'horizontal',
        //            left: 'right',
        top: 'bottom'
    },
    xAxis: {
        name: 'number of concurrent clients',
        nameLocation: 'middle',
        nameGap: 20,
        data: ["10", "100", "1000", "5000"]
    },
    yAxis: {
        name: 'total requests'
    },
    series: [
        {
            name: 'Xephon-K(Mem)',
            type: 'line',
            data: [12327, 21099, 31791, 12279]
        }, {
            name: 'Xephon-K(Cassandra)',
            type: 'line',
            data: [7931, 11336, 14590, 8703]
        },
        {
            name: 'KairosDB',
            type: 'line',
            data: [15561, 26154, 26939, 16506]
        },
        {
            name: 'InfluxDB',
            type: 'line',
            data: [118, 139, 131, 130]
        }
    ]
};

onlyChart.setOption(option);
// end of echarts

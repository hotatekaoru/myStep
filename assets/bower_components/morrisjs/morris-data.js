$(function() {

    Morris.Bar({
        element: 'morris-bar-chart',
        data: [{
            y: 'Coding',
            a: 100,
            b: 90
        }, {
            y: 'Training',
            a: 75,
            b: 65
        }, {
            y: 'Housework',
            a: 50,
            b: 40
        }, {
            y: 'Total',
            a: 120,
            b: 10
        }],
        xkey: 'y',
        ykeys: ['a', 'b'],
        labels: ['Kaoru', 'Yuri'],
        hideHover: 'auto',
        resize: true,
        barColors: ['rgb(31,92,183)','rgb(91,184,121)'],
    });
    
});

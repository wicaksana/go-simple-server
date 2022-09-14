function showAlertMessageSuccess() {
    $('#alertMessageSuccess').show('fade');
    setTimeout(function () {
        $('#alertMessageSuccess').hide('fade');
    }, 2000);
}

function showAlertMessageFailed() {
    $('#alertMessageFailed').show('fade');
    setTimeout(function () {
        $('#alertMessageFailed').hide('fade');
    }, 2000);
}

$(document).ready(function () {
    $('#formInfoLogMsg').submit(function (event) {
        var formInfoLogMsgData = {
            message: $("#textAreaInfoMsg").val(),
        };

        $.ajax({
            type: "POST",
            url: "info/new",
            data: JSON.stringify(formInfoLogMsgData),
            dataType: "json",
            encode: true,
        }).fail(function (data) {
            console.error(data);
            showAlertMessageFailed();
        }).done(function (data) {
            console.log(data);
            showAlertMessageSuccess();
        });
    });

    $('#formWarningLogMsg').submit(function (event) {
        var formWarningLogMsgData = {
            message: $("#textAreaWarningMsg").val(),
        };

        $.ajax({
            type: "POST",
            url: "warning/new",
            data: JSON.stringify(formWarningLogMsgData),
            dataType: "json",
            encode: true,
        }).fail(function (data) {
            console.error(data);
            showAlertMessageFailed();
        }).done(function (data) {
            console.log(data);
            showAlertMessageSuccess();
        });
    });


    $("#formErrorLogMsg").submit(function (event) {
        var formErrorLogMsgData = {
            message: $("#textAreaErrorMsg").val(),
        };

        $.ajax({
            type: "POST",
            url: "error/new",
            data: JSON.stringify(formErrorLogMsgData),
            dataType: "json",
            encode: true,
        }).fail(function (data) {
            console.error(data);
            showAlertMessageFailed();
        }).done(function (data) {
            console.log(data);
            showAlertMessageSuccess();
        });
    });

    // Metrics Generator
    $("#formMetricsGenerator").submit(function() {
        var metricsData = {
            temperature: parseFloat($("#inputMetricsTemperature").val()),
            humidity: parseFloat($("#inputMetricsHumidity").val()),
            precipitation: parseFloat($("#inputMetricsPrecipitation").val())
        };

        $.ajax({
            type: "POST",
            url: "metrics/new",
            data: JSON.stringify(metricsData),
            dataType: "json",
            encode: true,
        }).fail(function (data) {
            console.error(data);
            $('#alertMetricsFailed').show('fade');
            setTimeout(function () {
                $('#alertMetricsFailed').hide('fade');
            }, 2000);
        }).done(function (data) {
            console.log(data);
            $('#alertMetricsSuccess').show('fade');
            setTimeout(function () {
                $('#alertMetricsSuccess').hide('fade');
            }, 2000);
        });
    });

    // metrics alert button
    $('#buttonAlertMetricsSuccess').click(function () {
       $('#alertMetricsSuccess').hide('fade'); 
    });

    // msg alert button
    $('#buttonAlertMessageSuccess').click(function () {
        $('#alertMessageSuccess').hide('fade'); 
     });
});
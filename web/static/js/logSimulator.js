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
        }).done(function (data) {
            console.log(data);
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
        }).done(function (data) {
            console.log(data);
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
        }).done(function (data) {
            console.log(data);
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
        }).done(function (data) {
            console.log(data);
        });
    });
});
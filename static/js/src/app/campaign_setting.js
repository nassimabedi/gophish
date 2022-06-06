$(document).ready(function () {
    console.log("jquery");
    $("#settingsForm").submit(function (e) {
        alert(" form submit")
        /*$.post("/campaign_setting", $(this).serialize())
            .done(function (data) {
                alert("post success");
                successFlash(data.message)
            })
            .fail(function (data) {
                alert("post fail");
               // errorFlash(data.responseJSON.message)
            })*/
            
            campaign_setting = {
                duration: $("#duration").val(),                
           }
           api.campaignSetting.post(campaign_setting)
            .success(function (data) {
                alert("api success");
                resolve()
                campaign = data
            })
            .error(function (data) {
                alert("api fail");
                $("#modal\\.flashes").empty().append("<div style=\"text-align:center\" class=\"alert alert-danger\">\
    <i class=\"fa fa-exclamation-circle\"></i> " + data.responseJSON.message + "</div>")
                Swal.close()
            })
        return false
    })
    //$("#imapForm").submit(function (e) {
    $("#savesettings").click(function() {
        alert(" save click")
        
        return false
    })
}(jQuery));

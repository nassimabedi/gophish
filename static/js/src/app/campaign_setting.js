$(document).ready(function () {
    console.log("jquery");
    function confirmDelete() {
        swal({
            title: "Are you sure?",
            text: "You will not be able to recover this imaginary file!",
            type: "warning",
            showCancelButton: true,
            confirmButtonColor: "#DD6B55",
            confirmButtonText: "Delete",
            cancelButtonText: "Cancel",
            closeOnConfirm: false,
            closeOnCancel: false
        },
        function(isConfirm) {
            if (isConfirm) {
                document.deleteForm.submit();
            }
        });
    }
    api.campaignsttt.get()
        .success(function (result) {
            $("#duration").val(result.duration)
        })
        .error(function () {
            errorFlash("Error fetching IMAP settings")
        })
    /*$("#settingsForm").submit(function (e) {

        Swal.fire({
            title: "Are you sure?",
            text: "This will schedule the campaign to be launched.",
            type: "question",
            animation: false,
            showCancelButton: true,
            confirmButtonText: "Launch",
            confirmButtonColor: "#428bca",
            reverseButtons: true,
            allowOutsideClick: false,
            showLoaderOnConfirm: true,
            preConfirm: function () {
                return new Promise(function (resolve, reject) {                       
                    campaign_setting = {
                        duration: parseInt($("#duration").val()),                
                    }

                    api.campaignsttt.post(campaign_setting)
                            .success(function (data) {
                                campaign = data
                            })
                            .error(function (data) {
                                $("#modal\\.flashes").empty().append("<div style=\"text-align:center\" class=\"alert alert-danger\">\
                    <i class=\"fa fa-exclamation-circle\"></i> " + data.responseJSON.message + "</div>")
                                Swal.close()
                            })                   
                })
            }
        }).then(function (result) {
            if (result.value){
                Swal.fire(
                    'Campaign Scheduled!',
                    'This campaign has been scheduled for launch!',
                    'success'
                );
            }
            $('button:contains("OK")').on('click', function () {
                window.location = "/campaign_setting/"
            })
        })

    })*/

    $("#settingsForm").submit(function (e) {


        /*e.preventDefault();
  var nm_unit = $("#duration").val();
  //var almtunit = $("#almtunit").val();
  var form = this;

  swal({
    title: "Are you sure?",
    type: "warning",
    showCancelButton: true,
    confirmButtonColor: "#DD6B55",
    confirmButtonText: "Yes!",
    cancelButtonText: "Cancel",
    closeOnConfirm: true
  }, function(isConfirm) {
    if (isConfirm) {
      form.submit();
    }
  });*/

        /*e.preventDefault();

  //var data = $(this).serialize();

  swal({
    title: "Confirm?",
    text: "Are you sure?",
    type: "warning",
    showCancelButton: true,
    confirmButtonColor: "#DD6B55",
    confirmButtonText: "Confirm",
    cancelButtonText: "Back",
    preConfirm: function () {
        return new Promise(function (resolve, reject) {
           
             campaign_setting = {
            duration: parseInt($("#duration").val()),                
        }

        api.campaignsttt.post(campaign_setting)
                .success(function (data) {
                    campaign = data
                })
                .error(function (data) {
                    $("#modal\\.flashes").empty().append("<div style=\"text-align:center\" class=\"alert alert-danger\">\
        <i class=\"fa fa-exclamation-circle\"></i> " + data.responseJSON.message + "</div>")
                    Swal.close()
                })   
        })
    }
    }


  ).then(
    function (isConfirm) {
      if (isConfirm) {
        console.log('CONFIRMED');
      }
    },
    function() {
       console.log('BACK');
    }
  );

  return false;*/
  var form = this;
       
        campaign_setting = {
            duration: parseInt($("#duration").val()),                
        }

        api.campaignsttt.post(campaign_setting)
                .success(function (data) {
                    campaign = data
                })
                .error(function (data) {
                    $("#modal\\.flashes").empty().append("<div style=\"text-align:center\" class=\"alert alert-danger\">\
        <i class=\"fa fa-exclamation-circle\"></i> " + data.responseJSON.message + "</div>")
                    Swal.close()
                })   
                
                

    
                /*Swal.fire({
                    title: 'Are you sure?',
                    text: "You won't be able to revert this!",
                    icon: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#3085d6',
                    cancelButtonColor: '#d33',
                    confirmButtonText: 'Yes, delete it!'
                  }).then((result) => {
                    if (result.isConfirmed) {
                        campaign_setting = {
                            duration: parseInt($("#duration").val()),                
                        }
                
                        api.campaignsttt.post(campaign_setting)
                                .success(function (data) {
                                    campaign = data
                                })
                                .error(function (data) {
                                    $("#modal\\.flashes").empty().append("<div style=\"text-align:center\" class=\"alert alert-danger\">\
                        <i class=\"fa fa-exclamation-circle\"></i> " + data.responseJSON.message + "</div>")
                                   // Swal.close()
                                })
                    }
                  })*/


    })    
      
        
}(jQuery));

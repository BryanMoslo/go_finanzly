$(() => {

  $('#expense-Paid').click(function() {
    if ($(this).is(':checked') === true) {
      $(this).attr('value', true);
    }else{
      $(this).attr('value', false);
    }
  });


  $(".pay-exit").click(function(){
    let expenseId = $(this).data("expense");
    let boardId = $(this).data("board");
    let checkboxValue = $(this).is(':checked');

    $.ajax({
      type: "PUT",
      url: `/paid_expense/${expenseId}`,
      data: {"BoardID": boardId, "Paid": checkboxValue, params: {"expense_id": expenseId}},
      success: function(){},
      error: function(){}
    });
  });
});

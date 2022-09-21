function submitData() {
  let name = document.getElementById("input-nama").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone-number").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("message").value;

  if (name == "") {
    return alert("Nama harus diisi!");
  } else if (email == "") {
    return alert("Email harus diisi!");
  } else if (phone == "") {
    return alert("Nomor Telephon harus diisi!");
  } else if (subject == "") {
    return alert("Subject harus diisi!");
  } else if (message == "") {
    return alert("Pesan harus diisi!");
  }

  let linkToEmail = document.createElement("a");
  linkToEmail.href = `mailto:$rick.roll@gmail.com?subject=${subject}&body=Hallo, saya ${name}, ${subject}, ${message}, silahkan kontak ke nomor ${phone}`;
  linkToEmail.click();
}

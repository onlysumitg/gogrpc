
https://developers.google.com/protocol-buffers/docs/proto3#options


// -----------------------------------------

Reserved message number/name

  message Person {
    reserved 12, 13, 9 to 11; // these field number are reserved and can not be use
    reserved "foo", "bar"; // these field names are reserved and can not be use
 ..
}


//------------------------------------------------

Nested Types

message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}

If you want to reuse this message type outside its parent message type, you refer to it as _Parent_._Type_:

message SomeOtherMessage {
  SearchResponse.Result result = 1;  // <<  use a field from Nested type
}
 

//------------------------------------------------

Oneof
If you have a message with many fields and where at most one field will be set at the same time, you can enforce this behavior and save memory by using the oneof feature.

message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}

You can add fields of any type, except map fields and repeated fields.

//--------------------------------------------------------------------------------


Maps

map<key_type, value_type> map_field = N;

.where the key_type can be any integral or string type (so, any scalar type except for floating point types and bytes). 

Note that enum is not a valid key_type. 

The value_type can be any type except another map.

Map fields cannot be repeated.


// --------------------------------------------------------------------

Options

https://developers.google.com/protocol-buffers/docs/proto3#options
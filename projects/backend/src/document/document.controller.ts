import { Body, Controller, Get, Param, Post, Put } from '@nestjs/common';
import { DocumentService } from './document.service';
import { UpdateDocDto, CreateDocDto } from './dto';
import { APIResponse } from 'common';
import { Document } from '../store';

@Controller('api/document')
export class DocumentController {
  constructor(private readonly docService: DocumentService) {}

  @Post()
  create(): APIResponse<CreateDocDto> {
    const content = '';
    const id = this.docService.create(content);
    const resp: APIResponse<CreateDocDto> = {
      status: 0,
      data: {
        id,
        content,
      },
    };
    return resp;
  }

  @Get(':id')
  query(@Param('id') id: string): APIResponse<Document> {
    const doc = this.docService.query(id);
    const resp: APIResponse<Document> = {
      status: 0,
    };
    if (!doc) {
      resp.status = 1;
      resp.message = 'not found';
      return resp;
    }
    resp.data = doc;
    return resp;
  }

  @Put(':id')
  update(
    @Param('id') id: string,
    @Body() updateDocDto: UpdateDocDto,
  ): APIResponse<string> {
    this.docService.update(id, updateDocDto.content);
    const resp: APIResponse<string> = {
      status: 0,
      message: 'ok',
    };
    return resp;
  }
}
